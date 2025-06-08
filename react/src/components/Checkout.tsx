import {useLocation, useNavigate} from "react-router-dom";
import type { Product } from "../types/types.ts";
import axios from "axios";

export function Checkout() {
  const location = useLocation();
  const selectedProducts = (location.state as Product[]) || [];
  const navigate = useNavigate();

  const totalPrice = selectedProducts.reduce((total, product) => total + product.Price, 0).toFixed(2);

  const pay = async () => {
    const response = await axios.post("http://localhost:8080/payments", {Amount: Number.parseFloat(totalPrice)})
    if (response.status === 201) {
      alert("Payment successful!");
      navigate("/");
    } else {
      alert("Payment failed. Please try again.");
    }
  }

  return (
    <div>
      <h1>Checkout</h1>
      {selectedProducts.length === 0 ? (
        <p>No products selected.</p>
      ) : (
        <ul>
          {selectedProducts.map((product) => (
            <li key={product.ID}>
              {product.Name} - ${product.Price}
            </li>
          ))}
        </ul>
      )}
      {
        selectedProducts.length > 0 && (
          <div>
            <p>
              Total: $
              {totalPrice}
            </p>
            <button style={{ backgroundColor: "whitesmoke", color: "black" }} onClick={pay}>Pay</button>
          </div>
        )
      }
    </div>
  );
}