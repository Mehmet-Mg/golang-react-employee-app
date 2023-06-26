import React from 'react'
import ReactDOM from 'react-dom/client'
import { RouterProvider, createBrowserRouter } from 'react-router-dom'
import Root from './routes/root'
import ErrorPage from './error-page'
import Employees from './routes/contacts'
import EmployeeList from './components/EmployeeList'
import EmployeeAdd from './routes/employeeAdd'

const router = createBrowserRouter([
  {
    path: "/",
    element: <Root />,
    errorElement: <ErrorPage />,
    children: [
      {
        path: "employees",
        element: <EmployeeList />
      },
      {
        path: "employees/:employeeId",
        element: <Employees />
      },
      {
        path: "/employees/add",
        element: <EmployeeAdd />
      }
    ]
  }
])

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
)
