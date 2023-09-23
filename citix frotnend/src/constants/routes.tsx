import { Home, Portal } from "@pages";
import { IRoute, Role } from "@ts/types";

export const ROUTES: IRoute[] = [
  {
    name: "Home",
    path: "/",
    component: <Home />,
    roles: [Role.CLIENT, Role.ADMIN],
    isPublic: true,
  },
  {
    name: "Portal",
    path: "/portal/:roomID",
    component: <Portal />,
    roles: [Role.CLIENT, Role.ADMIN],
    isPublic: true,
  },
];
