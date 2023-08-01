-- name: ListCategories :many
SELECT * FROM categories
ORDER BY sort_index;

-- name: GetCategoryByID :one
SELECT * FROM categories
WHERE id = @id;

-- name: GetParentCategoriesWithCount :many
SELECT c.*, (SELECT count(*) FROM categories WHERE child_of = c.id and is_active = true) as child_count FROM categories c
WHERE c.child_of is null and is_active = true
ORDER BY c.sort_index;

-- name: GetCategoriesWithCountByParentID :many
SELECT c.*, (SELECT count(*) FROM categories WHERE child_of = c.id and is_active = true) as child_count FROM categories c
WHERE c.child_of = @parent_id and is_active = true
ORDER BY c.sort_index;

-- name: CreateCategory :one
INSERT INTO categories (id,
                        child_of,
                        category_name,
                        image_url,
                        is_active,
                        sort_index)
    VALUES (
            @id,
            @child_of,
            @category_name,
            @image_url,
            @is_active,
            @sort_index) RETURNING *;

-- name: UpdateCategory :one
UPDATE categories
    SET child_of = @child_of,
        category_name = @category_name,
        image_url = @image_url,
        is_active = @is_active,
        sort_index = @sort_index
    WHERE id = @id RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = @id;
