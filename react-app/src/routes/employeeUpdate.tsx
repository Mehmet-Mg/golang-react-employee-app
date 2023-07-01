import { Button, Col, Input, Row } from "antd";
import { useNavigate, useParams } from "react-router-dom";
import { Employee } from "../models/Employee";
import { useEffect, useState } from "react";
import { useEmployee } from "../useEmployee";

export default function Employees() {
  const {employeeId} = useParams();
  const navigate = useNavigate();

  const {getEmployee, updateEmployee} = useEmployee();

  const [employee, setEmployee] = useState<Employee>({
    id: 0,
    firstName: "",
    lastName: "",
    salary: 0,
    description: "",
  });

  useEffect(() =>{
    if(employeeId) {
      getEmployee(parseInt(employeeId)).then(data => setEmployee(data))
    }
  }, [])

  const handleChange = (event: any) => {
    setEmployee({
      ...employee,
      [event.target.name]: event.target.value
    })
  }

  const handleUpdateClick = async (employee: Employee) => {
      await updateEmployee(employee);
      navigate("/employees")
  }

  return (
    <>
      <Row gutter={[8, 16]}>
        <Col span={24}>
          <Input placeholder="First Name" name="firstName" value={employee?.firstName} onChange={handleChange}/>
        </Col>
        <Col span={24}>
          <Input placeholder="Last Name"  name="lastName" value={employee?.lastName} onChange={handleChange}/>
        </Col>
        <Col span={24}>
          <Input placeholder="Salary"name="salary" inputMode="decimal" value={employee?.salary}  onChange={handleChange}/>
        </Col>
        <Col span={24}>
          <Input placeholder="Description" name="description" inputMode="text" value={employee?.description}  onChange={handleChange}/>
        </Col>
        <Col span={12}>
          <Button block type="primary" onClick={() => handleUpdateClick(employee)}>
            Update
          </Button>
        </Col>
        <Col span={12}>
          <Button block onClick={() => navigate("/employees")}>Cancel</Button>
        </Col>
      </Row>
    </>
  );
}