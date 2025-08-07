package pgsql

import (
	"context"
	"database/sql"
	"fmt"
	"prakarsa-app/domain"
	"prakarsa-app/transport/request"
	"prakarsa-app/transport/response"
	"strings"
)

type pgsqlPartnerRepository struct {
	db *sql.DB
}

// NewPgsqlPartnerRepository will create new an todoRepository object representation of PartnerRepository interface
func NewPgsqlPartnerRepository(db *sql.DB) *pgsqlPartnerRepository {
	return &pgsqlPartnerRepository{
		db: db,
	}
}

func (r *pgsqlPartnerRepository) Create(ctx context.Context, partner *domain.Partner) (err error) {
	query := `INSERT INTO partner_types (id, name, description, is_active, created_by, created_at) VALUES ($1, $2, $3, $4, $5, $6)`
	if _, err = r.db.ExecContext(ctx, query, partner.ID, partner.Name, partner.Description, partner.IsActive, partner.CreatedBy, partner.CreatedAt); err != nil {
		return err
	}

	return
}

func (r *pgsqlPartnerRepository) Update(ctx context.Context, partner *domain.Partner) (err error) {
	// Build dynamic SET clauses from Partner struct
	sets := []string{}
	args := []interface{}{}
	idx := 1

	if partner.Name != "" {
		sets = append(sets, fmt.Sprintf("name = $%d", idx))
		args = append(args, partner.Name)
		idx++
	}

	if partner.Description != "" {
		sets = append(sets, fmt.Sprintf("description = $%d", idx))
		args = append(args, partner.Description)
		idx++
	}

	if partner.IsActive != nil {
		sets = append(sets, fmt.Sprintf("is_active = $%d", idx))
		args = append(args, partner.IsActive)
		idx++
	}

	// kalau ada sesuatu untuk di‐update, commit ke SQL
	if len(sets) > 0 {
		// Update stamp
		sets = append(sets, fmt.Sprintf("updated_at = $%d", idx))
		args = append(args, partner.UpdatedAt)
		idx++

		sets = append(sets, fmt.Sprintf("updated_by = $%d", idx))
		args = append(args, partner.UpdatedBy)
		idx++

		// tambahkan WHERE id = $idx
		args = append(args, partner.ID)
		query := fmt.Sprintf(
			"UPDATE partner_types SET %s WHERE id = $%d",
			strings.Join(sets, ", "),
			idx,
		)

		if _, err = r.db.ExecContext(ctx, query, args...); err != nil {
			return
		}
	}

	return
}

func (r *pgsqlPartnerRepository) Delete(ctx context.Context, partner *domain.Partner) (rowsAffected int64, err error) {
	query := "DELETE FROM partner_types WHERE id = $1"
	res, err := r.db.ExecContext(ctx, query, partner.ID)
	if err != nil {
		return
	}

	rowsAffected, err = res.RowsAffected()
	if err != nil {
		return
	}

	return
}

func (r *pgsqlPartnerRepository) GetList(ctx context.Context, request *request.GetListPartnerReq) (res []response.GetListPartnerRes, meta response.MetaRes, err error) {
	// 1. Build WHERE clauses
	wheres := []string{}
	args := []interface{}{}
	idx := 1

	if request.Name != "" {
		wheres = append(wheres, fmt.Sprintf("name ILIKE $%d", idx))
		args = append(args, "%"+request.Name+"%")
		idx++
	}

	if request.IsActive != nil {
		wheres = append(wheres, fmt.Sprintf("is_active = $%d", idx))
		args = append(args, request.IsActive)
		idx++
	}

	whereSQL := ""
	if len(wheres) > 0 {
		whereSQL = "WHERE " + strings.Join(wheres, " AND ")
	}

	// --- 2. Hitung totalCount dulu (tanpa LIMIT/OFFSET) ---
	countQuery := fmt.Sprintf(
		"SELECT COUNT(*) FROM partner_types %s",
		whereSQL,
	)
	if err = r.db.QueryRowContext(ctx, countQuery, args...).Scan(&meta.TotalData); err != nil {
		return nil, meta, err
	}

	// 2. Calculate LIMIT & OFFSET
	perPage := request.PerPage
	if perPage <= 0 {
		perPage = 10
	}
	page := request.Page
	if page <= 0 {
		page = 1
	}

	// total pages = ceil(total / perPage)
	meta.Page = page
	meta.PerPage = perPage
	meta.TotalPages = (meta.TotalData + perPage - 1) / perPage

	offset := (page - 1) * perPage

	// add LIMIT & OFFSET to args
	args = append(args, perPage, offset)
	limitPos, offsetPos := idx, idx+1

	// 3. Final query
	query := fmt.Sprintf(`
        SELECT
            id, name, description,
            is_active
        FROM partner_types
        %s
        ORDER BY created_at DESC
        LIMIT $%d OFFSET $%d
    `, whereSQL, limitPos, offsetPos)

	// 4. Execute
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, meta, err
	}
	defer rows.Close()

	// 5. Scan results
	for rows.Next() {
		var item response.GetListPartnerRes

		if err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Description,
			&item.IsActive,
		); err != nil {
			return nil, meta, err
		}

		res = append(res, item)
	}
	if errRow := rows.Err(); errRow != nil {
		return nil, meta, errRow
	}

	return
}
