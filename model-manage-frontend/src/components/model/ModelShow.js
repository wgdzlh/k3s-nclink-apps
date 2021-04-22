import {
  Show,
  SimpleShowLayout,
  TextField,
  DateField,
  NumberField,
} from "react-admin";

const ModelShow = (props) => {
  return (
    <Show title="Show Model" {...props}>
      <SimpleShowLayout>
        <TextField source="id" />
        <TextField source="name" />
        <DateField showTime source="created_at" />
        <DateField showTime source="updated_at" />
        <NumberField source="used" />
        <TextField source="def" />
      </SimpleShowLayout>
    </Show>
  );
};

export default ModelShow;
