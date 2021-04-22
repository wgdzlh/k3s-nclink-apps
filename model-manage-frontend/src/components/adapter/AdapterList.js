import {
  List,
  Datagrid,
  TextField,
  DateField,
  EditButton,
  DeleteButton,
} from "react-admin";

const AdapterList = (props) => {
  return (
    <List {...props}>
      <Datagrid rowClick="show">
        <TextField source="name" />
        {/* <Button label="RN"  color="secondary" /> */}
        <DateField source="created_at" />
        <DateField source="updated_at" />
        <TextField source="dev_id" />
        <TextField source="model_name" />
        <EditButton label="" />
        <DeleteButton label="" />
      </Datagrid>
    </List>
  );
};

export default AdapterList;
