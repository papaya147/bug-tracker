import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import ErrorModel from "../../error/ErrorModel";
import Bug from "../../model/Bug";
import getBug from "../../requests/bug/getBug";
import ErrorList from "../../error/ErrorList";
import closeBug from "../../requests/bug/closeBug";

const ResolveBug = () => {
  const { id } = useParams();
  const [bug, setBug] = useState<Bug | null>(null);
  const [remark, setRemark] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<ErrorModel | null>(null);
  const navigate = useNavigate();

  useEffect(() => {
    getBug(id ?? "").then((data) => {
      if (data.error.errors) navigate(-1);
      setBug(data.bug);
    });
  }, [id, navigate]);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    closeBug(id ?? "", remark).then((data) => {
      if (data.error.errors) setError(data.error);
      else navigate(-1);
    });
  };

  return (
    <div className="form">
      <h2>Resolve Bug</h2>
      <form onSubmit={handleSubmit}>
        <label>Name</label>
        <input value={bug?.name} disabled />
        <label>Description</label>
        <textarea value={bug?.description} disabled />
        <label>Remarks</label>
        <textarea value={remark} onChange={(e) => setRemark(e.target.value)} />
        {!isLoading && <button type="submit">Resolve</button>}
        {isLoading && <button disabled>Loading...</button>}
        {error && <ErrorList messages={error} />}
      </form>
      <button onClick={() => navigate(-1)}>Go Back</button>
    </div>
  );
};

export default ResolveBug;
