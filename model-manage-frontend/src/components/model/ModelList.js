import {
  List,
  Datagrid,
  TextField,
  DateField,
  NumberField,
  FunctionField,
  EditButton,
  CloneButton,
  DeleteButton,
} from "react-admin";

const ModelList = (props) => {
  return (
    <List {...props}>
      <Datagrid rowClick="show">
        <TextField source="id" color="primary" />
        {/* <Button label="RN"  color="secondary" /> */}
        <DateField source="created_at" />
        <DateField source="updated_at" />
        <NumberField source="used" />
        <FunctionField
          label="Def"
          render={(record) =>
            record.def.slice(0, 50) + (record.def.length > 50 ? " ..." : "")
          }
        />
        <EditButton label="" />
        <CloneButton label="" />
        <DeleteButton label="" />
      </Datagrid>
    </List>
  );
};

export default ModelList;
