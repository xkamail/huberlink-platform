import { VirtualCategoryEnum } from '@/lib/types'

export const virtualIcon = (category: VirtualCategoryEnum) => {
  let img = require(`@/assets/images/remote-control.png`)
  if (category === VirtualCategoryEnum.AirConditioner) {
    img = require(`@/assets/images/air-conditioner.png`)
  }
  if (category === VirtualCategoryEnum.TV) {
    img = require(`@/assets/images/tv.png`)
  }
  if (category === VirtualCategoryEnum.Fan) {
    img = require(`@/assets/images/fan.png`)
  }
  if (category === VirtualCategoryEnum.Speaker) {
    img = require(`@/assets/images/speaker.png`)
  }
  return img
}
