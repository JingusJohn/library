import { redirect } from "@sveltejs/kit";

export const load = ({ locals, fetch }: any) => {
  const getBooks = async () => {
    const booksRes = await fetch('http://localhost:8080/get-books');
    const booksData = await booksRes.json();
    return booksData;
  }

  if (!locals.pb.authStore.isValid) {
    throw redirect(303, "/");
  } else {
    return {
      books: getBooks()
    }
  }
}
