import { Avatar, Button, List, Modal, Skeleton } from "antd";
import { useNavigate } from "react-router-dom";
import { useEmployee } from "../useEmployee";
import { useEffect, useRef, useState } from "react";

export default function EmployeeList() {
  const navigate = useNavigate();
  const { listEmployee, getEmployeeData, isLoading, deleteEmployee } =
    useEmployee();
  const [openModal, setOpenModal] = useState<boolean>(false);
  const employeeId = useRef<string>("");

  useEffect(() => {
    getEmployeeData().then();
  }, []);

  const handleDeleteClick = async (employeeId: string) => {
    await deleteEmployee(employeeId);
    setOpenModal(false);
  };

  return (
    <>
    <Modal
    title="Delete Employee"
    open={openModal}
    onOk={() => handleDeleteClick(employeeId.current)}
    onCancel={() => setOpenModal(false)}
  >
    <p>Are you sure delete employee in list?</p>
  </Modal>
    <List
      className="demo-loadmore-list"
      loading={isLoading}
      itemLayout="horizontal"
      dataSource={listEmployee}
      renderItem={(employee) => (
        <List.Item
          actions={[
            <Button
              type="primary"
              onClick={() => {
                navigate(`/employees/${employee.id}`);
              }}
            >
              Edit
            </Button>,
            <Button type="primary" danger onClick={() => {
              employeeId.current = employee.id
              setOpenModal(true)
            }}>
              Delete
            </Button>,
          ]}
        >
          <Skeleton avatar title={false} loading={isLoading} active>
            <List.Item.Meta
              avatar={<Avatar />}
              title={
                <a href="https://ant.design">
                  {employee.firstName} {employee.lastName}
                </a>
              }
              description="Ant Design, a design language for background applications, is refined by Ant UED Team"
            />
            <div>$ {employee.salary}</div>
          </Skeleton>
        </List.Item>
      )}
    />
    </>
  );
}