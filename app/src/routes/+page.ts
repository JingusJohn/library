
export const load = async ({ fetch }: any) => {
  const getBookCount = async () => {
    const bookCountRes = await fetch("http://localhost:8080/get-book-count");
    const bookCountData = await bookCountRes.json();
    return bookCountData.count;
  }


  return {
    bookCount: getBookCount()
  }
}
