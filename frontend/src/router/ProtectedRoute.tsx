import { Navigate, Outlet } from "react-router-dom";
import { useAuth } from "../context/AuthContext"
import LoadingSpinner from "../ui/LoadingSpinner";

export function GuestProtectedRoute() {
    const { user, loading } = useAuth();
    if (loading) return <LoadingSpinner />
    if (user) return <Navigate to="/" />

    return <Outlet />
}

export function AdminProtectedRoute() {
    const { user, loading } = useAuth();
    if (loading) return <LoadingSpinner />
    if (user?.role_id !== 1 || !user) return <Navigate to="/" />

    return <Outlet />
}

export function PPKProtectedRoute() {
    const { user, loading } = useAuth();
    if (loading) return <LoadingSpinner />
    if (user?.role_id !== 2 || !user) return <Navigate to="/" />

    return <Outlet />
}

export function PokjaProtectedRoute() {
    const { user, loading } = useAuth();
    if (loading) return <LoadingSpinner />
    if (user?.role_id !== 3 || !user) return <Navigate to="/" />

    return <Outlet />
}

export function KepalaProtectedRoute() {
    const { user, loading } = useAuth();
    if (loading) return <LoadingSpinner />
    if ((user?.role_id !== 4 && user?.role_id !== 5) || !user) return <Navigate to="/" />

    return <Outlet />
}
