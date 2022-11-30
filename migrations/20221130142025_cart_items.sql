-- +goose Up
-- +goose StatementBegin
CREATE TABLE cart_items(
    user_id uuid NOT NULL,
    product_id uuid NOT NULL,
    quantity INT NOT NULL 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cart_items;
-- +goose StatementEnd
