import {
  List,
  Datagrid,
  TextField,
  DateField,
  NumberField,
  EditButton,
  DeleteButton,
} from "react-admin";

const ModelList = (props) => {
  return (
    <List {...props}>
      <Datagrid rowClick="show">
        <TextField source="name" />
        {/* <Button label="RN"  color="secondary" /> */}
        <DateField source="created_at" />
        <DateField source="updated_at" />
        <NumberField source="used" />
        <TextField source="def" />
        <EditButton label="" />
        <DeleteButton label="" />
      </Datagrid>
    </List>
  );
};

export default ModelList;
