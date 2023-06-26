import { Avatar, Button, List, Skeleton } from 'antd';
import { useNavigate } from 'react-router-dom';
import { useEmployee } from '../useEmployee';
import { useEffect } from 'react';

export default function EmployeeList() {
    const navigate = useNavigate();

    const {listEmployee, getEmployeeData, isLoading, deleteEmployee} = useEmployee();

    useEffect(() =>  {
      getEmployeeData().then();
    }, [])

  return (
    <List
    className="demo-loadmore-list"
    loading={isLoading}
    itemLayout="horizontal"
    dataSource={listEmployee}
    renderItem={(employee) => (
      <List.Item
        actions={[
          <Button type="primary" onClick={() => {
            navigate(`/employees/${employee.id}`)
          }}>Edit</Button>,
          <Button type="primary" danger onClick={async () => {
            await deleteEmployee(employee.id)
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
  )
}
