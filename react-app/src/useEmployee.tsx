import { useState } from "react";
import { Employee } from "./models/Employee";
import axios from "axios";
import { notification } from "antd";

const openNotification = (title: string, message: string, ) => {
    notification.open({
      message: title,
      description: message,
        onClick: () => {
        console.log('Notification Clicked!');
      },
    });
  };

export function useEmployee() { 
    const [listEmployee, setEmployeeList] = useState<Employee[]>([]);
    const [selectedEmployee, setSelectedEmpolyee] = useState<Employee>({
        id: "0",
        firstName: "",
        lastName: "",
        salary: 0
    });
    const [isLoading, setIsLoading] = useState<boolean>(false);

    const getEmployeeData = async () => {
        setIsLoading(true);
        const {data} = await axios.get("http://localhost:8080/employees")
        setIsLoading(false);
        setEmployeeList(data)
    }

    const getEmployee = async (employeeId: number): Promise<Employee> => {
        setIsLoading(true);
        const {data} = await axios.get(`http://localhost:8080/employees/${employeeId}`)
        setIsLoading(false);
        return data;
    }

    const updateEmployee = async (employee: Employee) => {
        employee.salary = parseFloat(employee.salary.toString())
        await axios.put(`http://localhost:8080/employees/${employee.id}`, employee)
    }

    const addEmployee = async (employee: Employee) => {
            employee.salary = parseFloat(employee.salary.toString())
            const {status, data} = await axios.post(`http://127.0.0.1:8080/employees`, employee)
            debugger;
    }

    const deleteEmployee = async (employeeId: string) => {
        await axios.delete(`http://localhost:8080/employees/${employeeId}`)
            .then(resp => {
                if(resp.status === 200) {
                    setEmployeeList(listEmployee.filter(t => t.id !== employeeId))
                }
            })
    }

    return {
      listEmployee,
      selectedEmployee,
      isLoading,
      setSelectedEmpolyee,
      getEmployeeData,
      getEmployee,
      updateEmployee,
      addEmployee,
      deleteEmployee,
    };
}