import {
  Edit,
  SimpleForm,
  TextInput,
  ReferenceInput,
  SelectInput,
} from "react-admin";

const AdapterEdit = (props) => {
  return (
    <Edit title="Edit Adapter" {...props}>
      <SimpleForm>
        <TextInput disabled source="id" />
        <TextInput source="dev_id" />
        <ReferenceInput source="model_id" reference="models">
          <SelectInput optionText="id" />
        </ReferenceInput>
      </SimpleForm>
    </Edit>
  );
};

export default AdapterEdit;
