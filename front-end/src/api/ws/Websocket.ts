export const rootWS = 'ws://localhost:8030/ws'

export const matchmakingUrl = `${rootWS}/match`
export const gameplayUrl = `${rootWS}/game`

export const matchmakingCodes = {
  timedOut: '104',
  stop: '103',
  found: '102',
  search: '101',
}

export const gameplayCodes = {
  connected: '200',
  start: '201',
  img: '202',
  chat: '203',
  wait: '204',
  updateImg: '205',
  updateChat: '206',
  timeUpdate: '207',
  endTime: '208',
  endWin: '209',
}
