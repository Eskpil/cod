import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import React from "react";
import { createBrowserRouter, RouterProvider, Route } from "react-router-dom";
import { DashboardPage } from "./pages/dashboard";
import { MachinePage } from "./pages/machines/machine";

interface Props {}

const router = createBrowserRouter([
  {
    path: "/",
    element: <DashboardPage />,
  },
  {
    path: "/test",
    element: <div className="h-screen bg-pink-300 text-black">Testpage</div>,
  },
  {
    path: "/machines/:id/:page",
    element: <MachinePage />,
  },
]);

const App: React.FC<Props> = () => {
  const queryClient = new QueryClient();

  return (
    <QueryClientProvider client={queryClient}>
      <RouterProvider router={router} />
    </QueryClientProvider>
  );
};

export default App;
