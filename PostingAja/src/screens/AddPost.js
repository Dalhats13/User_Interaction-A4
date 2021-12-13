import React, { Component } from 'react';
import { Image, Pressable, View, StyleSheet, TextInput, ScrollView } from 'react-native';
import { Text, Button, Input } from 'react-native-ui-kitten';
import * as ImagePicker from 'react-native-image-picker';
//import { withFirebaseHOC } from '../utils'

class AddPost extends Component {
  state = { image: null, title: '', description: '' };

  onChangeTitle = title => {
    this.setState({ title });
  };
  onChangeDescription = description => {
    this.setState({ description });
  };

  onSubmit = async () => {
    try {
      const post = {
        photo: this.state.image,
        title: this.state.title,
        description: this.state.description,
      };

      this.setState({
        image: null,
        title: '',
        description: '',
      });
    } catch (e) {
      console.error(e);
      const options = {
        noData: true,
        mediaType: 'photo',
      };
    }
  };

  selectImage = () => {
    const options = {
      noData: true,
    };
    ImagePicker.launchImageLibrary(options, response => {
      if (response.didCancel) {
        console.log('User cancelled image picker');
      } else if (response.error) {
        console.log('ImagePicker Error: ', response.error);
      } else if (response.customButton) {
        console.log('User tapped custom button: ', response.customButton);
      } else {
        if (response && response.hasOwnProperty('assets')) {
          const source = response.assets[0];
          this.setState({
            image: source,
          });
        }
      }
    });
  };

  render() {
    return (
      <ScrollView>
      <><View>
        <View>
          <View style={{ flexDirection: 'row', marginTop: 10 }}>
            <View style={{ flex: 1 }}>
              <Pressable
                status="success"
                style={{
                  width: 100,
                  height: 40,
                  marginBottom: 20,
                  backgroundColor: '#0C8EFF',
                  justifyContent: 'center',
                  alignItems: 'center',
                  borderRadius: 100 / 2,
                  // justifyContent: 'flex-start',
                }}
                onPress={this.onSubmit}
                disabled={this.state.image && this.state.title && this.state.description
                  ? false
                  : true}>
                <Text style={{ color: '#FFFFFF' }}>Back</Text>
              </Pressable>
            </View>
            <View style={{ flexDirection: 'column' }}>
            <View style={{ flex: 1 }}>
              <Pressable
                status="success"
                style={{
                  width: 100,
                  height: 40,
                  marginBottom: 20,
                  backgroundColor: '#0C8EFF',
                  justifyContent: 'center',
                  alignItems: 'center',
                  borderRadius: 100 / 2,
                  // justifyContent: 'flex-end',
                }}
                onPress={this.onSubmit}
                disabled={this.state.image && this.state.description
                  ? false
                  : true}>
                <Text style={{ color: '#FFFFFF' }}>Add post</Text>
              </Pressable>
            </View>
            </View>
          </View>
        </View>
        <View style={{ alignItems: 'center' }}>
          {this.state.image ? (
          <Image
            source={
              this.state.image
            }
            style={{
              width: 350,
              height: 400,
            }} />
           ) : (
      <Button
        onPress={this.selectImage}
        style={{
          alignItems: 'center',
          padding: 10,
          margin: 30,
        }}>
        Add an image
      </Button>
    )} 
        </View>
        <View style={styles.textAreaContainer}>
            {/* <Text category="h4">Post Details</Text> */}

            <TextInput
              placeholder="Enter caption"
              style={styles.textArea}
              underlineColorAndroid="transparent"
              placeholder="Write Caption..."
              placeholderTextColor="grey"
              numberOfLines={10}
              multiline={true}
              value={this.state.description}
              onChangeText={description =>
                this.onChangeDescription(description)} />
        </View>
      </View></>
      </ScrollView>

    );
  }
}

const styles = StyleSheet.create({
  textAreaContainer: {
    borderColor: 'transparent',
    borderWidth: 1,
    padding: 5,
    alignItems: 'flex-start',
  },
  textArea: {
    height: 150,
    marginBottom: 30,
    justifyContent: 'flex-start',
  },
});

export default AddPost;
