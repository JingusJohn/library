import { redirect } from "@sveltejs/kit";

export const load = ({ locals }) => {
  if (locals.pb.authStore.isValid) {
    throw redirect(303, "/");
  }
}

export const actions = {
  // default: async({ request }: any) => {
    // const formData = await request.formData();
    // console.log(formData);
  // },
  register: async ({ locals, request }: any) => {
    const formData = await request.formData();
    const data = Object.fromEntries([...formData]);
    console.log(data);


    try {
      console.log("logging info", data.email, typeof data.email)

      const newUser = await locals.pb.users.create(data);

      const { token, user } =await locals.pb.users.authViaEmail(data.email, data.password);

      const updatedProfile = await locals.pb.records.update('profiles', user.profile.id, { name: data.name });

      locals.pb.authStore.clear();
    } catch (err) {
      console.log("error:", err);
      return {
        error: true,
        message: err
      }
    }

    throw redirect(303, "/login");

    // try {
      // const newUser = await locals.pb.users.create(data);

      // const { token, user } = await locals.pb.users.authViaEmail(data.email, data.password);
      
      // const updatedProfile = await locals.pb.records.update('profiles', user.profile.id, { name: data.name });

      // locals.pb.authStore.clear()
    // } catch (err) {
      // console.log('Error:', err);
      // return  {
        // error: true,
        // message: err
      // };
    // }

    // throw redirect(303, '/login');
  }
}
