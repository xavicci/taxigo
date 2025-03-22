import { login } from '../services/grpc';

// En tu componente:
const handleLogin = async (email, password) => {
  try {
    const response = await login(email, password);
    // Manejar la respuesta exitosa
    console.log('Token:', response.token);
    console.log('User:', response.user);
  } catch (error) {
    // Manejar el error
    console.error('Error:', error);
  }
};