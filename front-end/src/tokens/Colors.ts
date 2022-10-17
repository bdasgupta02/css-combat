export const percentageToHex = (percentage: number) => {
  const decimal = Math.round((percentage * 255) / 100)
  if (percentage < 7) {
    return '0' + decimal.toString(16).toUpperCase()
  } else {
    return decimal.toString(16).toUpperCase()
  }
}

const Colors = {
  Transparent: '#00000000',
  White: '#FFFFFF',
  Background: '#262A35',
  Logo: '#F53F4F',
  Input: '#E8E8E8',
  InputFocused: '#5B8FDD',
  InputText: '#242424',
  InputTextFocused: '#FFFFFF',
  InputBorderHover: `#FFFFFF${percentageToHex(23)}`,
  InputBorderFocused: `#5B8FDD${percentageToHex(60)}`,
  SignInFrontBox: `#364966${percentageToHex(21)}`,
  SignInBackBox: `#000000${percentageToHex(17)}`,
  SignInDetails: `#FFFFFF${percentageToHex(12)}`,
  Button: `#F53F4F${percentageToHex(53)}`,
  ButtonHover: `#F53F4F${percentageToHex(77)}`,
  ButtonRoot: `#F53F4F`,
  DarkBg: `#000000${percentageToHex(17)}`,
  Blue: '#3E67A5',
  SectionHeading: `#FFFFFF${percentageToHex(30)}`
}

export default Colors
