
export const load = ({ fetch, params }: any) => {

  const fetchBook = async (id: string) => {
    const res = await fetch(`http://localhost:8080/get-book/${id}`);
    const data = await res.json();
    return data;
  }

  return {
    book: fetchBook(params.bookId)
  };
}
