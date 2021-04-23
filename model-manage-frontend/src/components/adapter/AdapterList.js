import {
  List,
  Datagrid,
  TextField,
  DateField,
  EditButton,
  CloneButton,
  DeleteButton,
} from "react-admin";

const AdapterList = (props) => {
  return (
    <List {...props}>
      <Datagrid rowClick="show">
        <TextField source="id" />
        {/* <Button label="RN"  color="secondary" /> */}
        <DateField source="created_at" />
        <DateField source="updated_at" />
        <TextField source="dev_id" />
        <TextField source="model_id" />
        <EditButton label="" />
        <CloneButton label="" />
        <DeleteButton label="" />
      </Datagrid>
    </List>
  );
};

export default AdapterList;
