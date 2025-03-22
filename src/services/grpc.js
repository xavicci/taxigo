import { NativeModules } from 'react-native';
const { GrpcModule } = NativeModules;

export const login = (email, password) => {
  return new Promise((resolve, reject) => {
    GrpcModule.login({ email, password })
      .then(response => resolve(response))
      .catch(error => reject(error));
  });
};

export const register = (email, password, name, phone) => {
  return new Promise((resolve, reject) => {
    GrpcModule.register({ email, password, name, phone })
      .then(response => resolve(response))
      .catch(error => reject(error));
  });
}; 