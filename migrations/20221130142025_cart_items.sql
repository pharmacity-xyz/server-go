-- +goose Up
-- +goose StatementBegin
CREATE TABLE cart_items(
    user_id uuid NOT NULL REFERENCES users (user_id),
    product_id uuid NOT NULL REFERENCES products (product_id),
    quantity INT NOT NULL 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cart_items;
-- +goose StatementEnd
