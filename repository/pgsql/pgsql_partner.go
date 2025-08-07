package pgsql

import (
	"context"
	"database/sql"
	"fmt"
	"prakarsa-app/domain"
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
