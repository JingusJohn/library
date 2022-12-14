import { redirect } from "@sveltejs/kit";

export const load = ({ locals }: any) => {
  if (locals.pb.authStore.isValid) {
    throw redirect(303, "/");
  }
}


export const actions = {
  login: async ({ request, locals }: any) => {
    const formData = await request.formData();
    const data = Object.fromEntries([...formData]);

    try {
      const { token, user } = await locals.pb.users.authViaEmail(data.email, data.password);
    } catch (err) {
      console.log("err:", err);
      return {
        error: true,
        email: data.email
      };
    }
    throw redirect(303, '/');
  }
}
