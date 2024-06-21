import React from 'react';
import {View, Text, FlatList, StyleSheet} from 'react-native';

const MessageList = ({messages}) => {
  return (
    <FlatList
      data={messages}
      keyExtractor={item => item.id.toString()}
      renderItem={({item}) => (
        <View style={styles.messageContainer}>
          <Text style={styles.message}>{item.message}</Text>
        </View>
      )}
      extraData={messages}
    />
  );
};

const styles = StyleSheet.create({
  messageContainer: {
    padding: 10,
    borderBottomWidth: 1,
    borderBottomColor: '#ccc',
  },
  message: {
    fontSize: 16,
  },
});

export default MessageList;
