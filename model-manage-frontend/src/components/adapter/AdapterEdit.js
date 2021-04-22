import { Edit, SimpleForm, TextInput } from "react-admin";

const AdapterEdit = (props) => {
  return (
    <Edit title="Edit Adapter" {...props}>
      <SimpleForm>
        <TextInput disabled source="name" />
        <TextInput source="dev_id" />
        <TextInput source="model_name" />
      </SimpleForm>
    </Edit>
  );
};

export default AdapterEdit;
