-- +goose Up
-- +goose StatementBegin
CREATE TABLE order_items(
    order_id uuid REFERENCES orders(order_id),
    product_id uuid REFERENCES products(product_id),
    quantity INT NOT NULL,
    total_price NUMERIC NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders_items;
-- +goose StatementEnd
