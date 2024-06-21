import axios from 'axios';

const API_URL = 'http://192.168.1.4:8080';

export const getMessages = async () => {
  try {
    const response = await axios.get(`${API_URL}/api/getMessage`);
    return response.data;
  } catch (error) {
    console.error('Error fetching messages:', error);
    throw error;
  }
};

export const createMessage = async message => {
  try {
    const response = await axios.post(`${API_URL}/api/createMessage`, {
      message,
    });
    return response.data;
  } catch (error) {
    console.error('Error creating message:', error);
    throw error;
  }
};
