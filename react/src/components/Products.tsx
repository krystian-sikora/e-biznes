import { useEffect, useState } from "react";
import axios from "axios";
import type { Product } from "../types/types.ts";
import "../index.css";
import { useNavigate } from "react-router-dom";

export function Products() {
  const [products, setProducts] = useState<Product[]>([]);
  const [selected, setSelected] = useState<Product[]>([]);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchProducts = async () => {
      const response = await axios.get("http://localhost:8080/products");
      setProducts(response.data);
    };
    fetchProducts();
  }, []);

  const toggleSelect = (product: Product) => {
    setSelected((prev) =>
      prev.find((p) => p.ID === product.ID)
        ? prev.filter((p) => p.ID !== product.ID)
        : [...prev, product]
    );
  };

  const goToCheckout = () => {
    navigate("/checkout", { state: selected });
  };

  return (
    <div>
      <div style={{ display: "flex", justifyContent: "flex-end", alignItems: "flex-start", width: "100%" }}>
        <button
          style={{ backgroundColor: "whitesmoke", color: "black" }}
          onClick={goToCheckout}
          disabled={selected.length === 0}
        >
          Checkout
        </button>
      </div>
      <div style={{ display: "flex", flexDirection: "column", alignItems: "center", justifyContent: "center" }}>
        <h1>Products</h1>
        <div style={{ display: "flex", flexWrap: "wrap", width: "75%", justifyContent: "center" }}>
          {products?.map((product) => (
            <div
              style={{
                width: 250,
                backgroundColor: selected.find((p) => p.ID === product.ID) ? "#b3e5fc" : "whitesmoke",
                color: "black",
                borderRadius: 15,
                margin: 10,
                padding: 10,
                border: selected.find((p) => p.ID === product.ID) ? "2px solid #0288d1" : "none"
              }}
              key={product.ID}
              onClick={() => toggleSelect(product)}
            >
              <h2>{product.Name}</h2>
              <p>Price: ${product.Price}</p>
              {product.Categories && product.Categories.length > 0 && (
                <p>Categories: {product.Categories.map((category) => category.Name).join(", ")}</p>
              )}
              <p>
                <input
                  type="checkbox"
                  checked={!!selected.find((p) => p.ID === product.ID)}
                  readOnly
                />{" "}
                Select
              </p>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}