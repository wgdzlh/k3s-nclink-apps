import {
  List,
  Datagrid,
  TextField,
  DateField,
  NumberField,
  EditButton,
  CloneButton,
  DeleteButton,
} from "react-admin";

const ModelList = (props) => {
  return (
    <List {...props}>
      <Datagrid rowClick="show">
        <TextField source="id" />
        {/* <Button label="RN"  color="secondary" /> */}
        <DateField source="created_at" />
        <DateField source="updated_at" />
        <NumberField source="used" />
        <TextField source="def" />
        <EditButton label="" />
        <CloneButton label="" />
        <DeleteButton label="" />
      </Datagrid>
    </List>
  );
};

export default ModelList;
