export default abstract class API {
  private static address = 'http://localhost:3000/api/';

  protected static async apiCall<T>(url: string, body: T): Promise<Response> {
    return fetch(`${API.address}${url}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(body),
    });
  }
}
