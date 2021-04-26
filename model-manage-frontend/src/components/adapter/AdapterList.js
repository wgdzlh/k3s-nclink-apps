import {
  List,
  Datagrid,
  TextField,
  DateField,
  UrlField,
  EditButton,
  CloneButton,
  DeleteButton,
  FunctionField,
} from "react-admin";

const AdapterList = (props) => {
  return (
    <List {...props}>
      <Datagrid rowClick="show">
        <TextField source="id" color="primary" />
        {/* <Button label="RN"  color="secondary" /> */}
        <DateField source="created_at" />
        <DateField source="updated_at" />
        <TextField source="dev_id" />
        <FunctionField
          source="model_id"
          render={(record) => (
            <UrlField
              source="model_id"
              color="secondary"
              href={`#/models/${record.model_id}/show`}
              onClick={(e) => e.stopPropagation()}
            />
          )}
        />
        <EditButton label="" />
        <CloneButton label="" />
        <DeleteButton label="" />
      </Datagrid>
    </List>
  );
};

export default AdapterList;
