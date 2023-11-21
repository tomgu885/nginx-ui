package model

func stateNormalize(state int8) int8 {
    if StateEnable == state {
        return StateEnable
    }

    return StateDisabled
}
