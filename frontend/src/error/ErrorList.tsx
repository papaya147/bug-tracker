import ErrorModel from "./ErrorModel";

const ErrorList = ({ messages }: { messages: ErrorModel }) => {
  let index = 0;
  return (
    <div className="error-list">
      {messages.errors.map((error, index) => {
        return (
          <div className="error-preview" key={String(index)}>
            {error.message}
          </div>
        );
      })}
    </div>
  );
};

export default ErrorList;
