import React, {useState, useEffect} from 'react';
import {View, TextInput, Button, StyleSheet} from 'react-native';
import MessageList from './src/components/MessageList';
import {getMessages, createMessage} from './src/services/api';
import {Pusher} from '@pusher/pusher-websocket-react-native';

const App = () => {
  const [messages, setMessages] = useState([]);
  const [input, setInput] = useState('');

  useEffect(() => {
    const fetchMessages = async () => {
      const fetchedMessages = await getMessages();
      setMessages(fetchedMessages);
    };

    fetchMessages();

    const setupPusher = async () => {
      const pusher = Pusher.getInstance();

      await pusher.init({
        apiKey: '3bf5567f779653cead1b',
        cluster: 'ap1',
      });

      await pusher.connect();

      const channel = await pusher.subscribe({
        channelName: 'chat-channel',
        onEvent: event => {
          const eventData = JSON.parse(event.data);
          setMessages(prevMessages => [...prevMessages, eventData]);
        },
      });

      return () => {
        channel.unsubscribe();
        pusher.disconnect();
      };
    };

    setupPusher().catch(error =>
      console.error('Error initializing Pusher:', error),
    );
  }, []);

  const sendMessage = async () => {
    if (input.trim()) {
      const newMessage = await createMessage(input);
      setMessages([...messages, newMessage]);
      setInput('');
    }
  };

  return (
    <View style={styles.container}>
      <MessageList messages={messages} />
      <TextInput
        style={styles.input}
        value={input}
        onChangeText={setInput}
        placeholder="Type your message"
      />
      <Button title="Send" onPress={sendMessage} />
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    padding: 20,
  },
  input: {
    borderWidth: 1,
    borderColor: '#ccc',
    padding: 10,
    marginBottom: 10,
  },
});

export default App;
