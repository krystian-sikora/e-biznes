import { Products } from './components/Products.tsx'
import {BrowserRouter, Route, Routes} from "react-router-dom";
import {Checkout} from "./components/Checkout.tsx";

function App() {

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Products />} />
        <Route path="/products" element={<Products />} />
        <Route path="/checkout" element={<Checkout />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App
