import {
  Show,
  SimpleShowLayout,
  TextField,
  DateField,
  FunctionField,
  UrlField,
} from "react-admin";

const AdapterShow = (props) => {
  return (
    <Show title="Show Adapter" {...props}>
      <SimpleShowLayout>
        <TextField source="id" />
        <DateField showTime source="created_at" />
        <DateField showTime source="updated_at" />
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
      </SimpleShowLayout>
    </Show>
  );
};

export default AdapterShow;
