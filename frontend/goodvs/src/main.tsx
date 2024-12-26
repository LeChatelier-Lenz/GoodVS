// import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import { RouterProvider,createBrowserRouter } from 'react-router-dom'
import ErrorPage from "./pages/error-page.tsx";
import Root from "./pages/root.tsx";
import Search from "./pages/search.tsx";
import User from "./pages/user.tsx";
import SignIn from "./pages/signin.tsx";
import SignUp from "./pages/signup.tsx";


const router = createBrowserRouter(
    [
        {
            path: '/',
            element: < Root />,
            errorElement: <ErrorPage />,
            children: [
                {
                    path: '/search',
                    element: <Search />,
                },
                {
                    path: '/user',
                    element: <User />,
                },
                // {
                //     path:'details:id',
                //     element: <Details />,
                // }
            ]
        },
        {
            path: '/signin',
            element: <SignIn />,
            errorElement: <ErrorPage />,
        },
        {
            path: '/signup',
            element: <SignUp />,
            errorElement: <ErrorPage />,
        },
    ]);

const root = createRoot(
    document.getElementById('root') as HTMLElement
);
root.render(
    <div>
        <RouterProvider router={router} />
    </div>
);
