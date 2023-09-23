import { Toaster } from "@components/ui/toaster";
import { ROUTES } from "@constants/routes";
import { AnimatePresence } from "framer-motion";
import { Route, Routes } from "react-router-dom";

const App: React.FC = () => {
  return (
    <div className="w-full h-screen">
      <AnimatePresence mode="wait">
        <Routes>
          {ROUTES.map((route) => (
            <Route
              key={route.path}
              path={route.path}
              element={route.component}
            />
          ))}
        </Routes>
        <Toaster />
      </AnimatePresence>
    </div>
  );
};

export default App;
