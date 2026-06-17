function Pagination({ page, limit, onPrevious, onNext, onLimitChange }) {
  return (
    <div className="pagination">
      <button className="btn secondary" onClick={onPrevious} disabled={page === 1}>
        Previous
      </button>

      <span>Page {page}</span>

      <button className="btn secondary" onClick={onNext}>
        Next
      </button>

      <select value={limit} onChange={(e) => onLimitChange(Number(e.target.value))}>
        <option value={5}>5 / page</option>
        <option value={10}>10 / page</option>
        <option value={20}>20 / page</option>
      </select>
    </div>
  );
}

export default Pagination;