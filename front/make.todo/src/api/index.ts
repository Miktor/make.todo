export default abstract class API {
  private static address = 'localhost:3000/api/';

  protected static async apiCall<T, R>(url: string, body: T): Promise<R> {
    return fetch(`${API.address}${url}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(body),
    }).then(data => data.json());
  }
}
