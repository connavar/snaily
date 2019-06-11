export class UserState {

  public static isLoggedIn(): boolean {
    return localStorage.getItem('login') === 'true';
  }

  public static login() {
    localStorage.setItem('login', 'true');
  }

  public static logout() {
    localStorage.setItem('login', 'false');
    this.setMyID('');
    this.setMyKey('');
  }

  public static setMyKey(v: string) {
    localStorage.setItem('key', v);
  }

  public static getMyKey(): string {
    return localStorage.getItem('key');
  }

  public static setMyID(v: string) {
    localStorage.setItem('id', v);
  }

  public static getMyID(): string {
    return localStorage.getItem('id');
  }


}
