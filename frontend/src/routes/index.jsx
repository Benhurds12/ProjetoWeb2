import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";

import Login from "../pages/Login";
import Bens from "../pages/Bens";
import Setores from "../pages/Setores";
import Fabricantes from "../pages/Fabricantes";
import Fornecedores from "../pages/Fornecedores";

function PrivateRoute({ children }) {
  const token = localStorage.getItem("@SCI:token");

  return token ? children : <Navigate to="/" />;
}

export default function RoutesApp() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Login />} />

        <Route
          path="/bens"
          element={
            <PrivateRoute>
              <Bens />
            </PrivateRoute>
          }
        />

        <Route
          path="/setores"
          element={
            <PrivateRoute>
              <Setores />
            </PrivateRoute>
          }
        />

        <Route
          path="/fabricantes"
          element={
            <PrivateRoute>
              <Fabricantes />
            </PrivateRoute>
          }
        />

        <Route
          path="/fornecedores"
          element={
            <PrivateRoute>
              <Fornecedores />
            </PrivateRoute>
          }
        />

        <Route path="*" element={<Navigate to="/" />} />
      </Routes>
    </BrowserRouter>
  );
}
