import { Products } from './components/Products.tsx'
import {BrowserRouter, Route, Routes} from "react-router-dom";
import {Checkout} from "./components/Checkout.tsx";
import {Hello} from "./components/Hello.tsx";

function App() {

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Products />} />
        <Route path="/products" element={<Products />} />
        <Route path="/checkout" element={<Checkout />} />
        <Route path="/hello" element={<Hello/>} />
      </Routes>
    </BrowserRouter>
  )
}

export default App
