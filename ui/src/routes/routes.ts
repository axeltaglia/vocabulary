import React from "react"
import Home from "../pages/Home";
import Categories from "../pages/Categories";

type RouteType = {
    key: string
    title: string
    path: string
    authenticated: boolean
    enabled: boolean
    component: React.ComponentType<any>
}

export const routes: RouteType[] = [
    {
        key: 'home-route',
        title: 'Home',
        path: '/',
        authenticated: false,
        enabled: true,
        component: Home
    },
    {
        key: 'category-route',
        title: 'Category',
        path: '/categories',
        authenticated: false,
        enabled: true,
        component: Categories
    }
]