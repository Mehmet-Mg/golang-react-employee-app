import { useState } from "react";
import { Employee } from "./models/Employee";
import axios from "axios";
import { message } from "antd";

export function useEmployee() { 
    const [listEmployee, setEmployeeList] = useState<Employee[]>([]);
    const [selectedEmployee, setSelectedEmpolyee] = useState<Employee>({
        id: 0,
        firstName: "",
        lastName: "",
        salary: 0,
        description: "",
    });
    const [isLoading, setIsLoading] = useState<boolean>(false);
    const [messageApi, contextHolder] = message.useMessage();

    const success = () => {
        messageApi.open({
          type: 'success',
          content: 'Succesfull',
        });
      };
    
      const error = () => {
        messageApi.open({
          type: 'error',
          content: 'This is an error message',
        });
      };
    
      const warning = () => {
        messageApi.open({
          type: 'warning',
          content: 'This is a warning message',
        });
      };


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
        try {
            employee.salary = parseFloat(employee.salary.toString())
            await axios.put(`http://localhost:8080/employees/${employee.id}`, employee)
            success()
        } catch(ex) {
            error()
        }
    }

    const addEmployee = async (employee: Employee) => {
        try {
            employee.salary = parseFloat(employee.salary.toString())
            await axios.post(`http://127.0.0.1:8080/employees`, employee)
            success()
        } catch (ex) {
            error()
        }
    }

    const deleteEmployee = async (employeeId: number) => {
        try {
            await axios.delete(`http://localhost:8080/employees/${employeeId}`)
            setEmployeeList(list => list.filter(emp => emp.id !== employeeId))
            success()
        } catch(ex) {
            error()
        }
    }

    return {
      listEmployee,
      selectedEmployee,
      isLoading,
      contextHolder,
      setSelectedEmpolyee,
      getEmployeeData,
      getEmployee,
      updateEmployee,
      addEmployee,
      deleteEmployee,
    };
}