import { Button, Col, Input, Row } from "antd";
import { Employee } from "../models/Employee";
import { useState } from "react";
import { useEmployee } from "../useEmployee";
import { useNavigate } from "react-router-dom";


export default function EmployeeAdd() {
  const navigate = useNavigate();
  const {addEmployee, contextHolder} = useEmployee();
  const [employee, setEmployee] = useState<Employee>({
    id: 0,
    firstName: "",
    lastName: "",
    salary: 0,
    description: "",
  });

  const handleChange = (event: any) => {
    setEmployee({
      ...employee,
      [event.target.name]: event.target.value
    })
  }

  const handleAddClick = async () => {
    await addEmployee(employee)
    navigate("/employees");
  }

  return (
    <>
    {contextHolder}
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
          <Input
            placeholder="Description"
            name="description"
            inputMode="text"
            value={employee?.description}
            onChange={handleChange}
          />
        </Col>
        <Col span={12}>
          <Button block type="primary" onClick={handleAddClick}>
            Add
          </Button>
        </Col>
        <Col span={12}>
          <Button block>Cancel</Button>
        </Col>
      </Row>
    </>
  );
}