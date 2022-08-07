import { writable } from 'svelte/store';

type Notification = {
    type: string;
    message: string;
    timeout: number;
    title: string;
}


type MapNotification = Map<string, Notification>;
export const notifications = writable<MapNotification>(new Map());

export const push = (title: string, message: string, messageType: string = 'danger', duration: number = 5000) => {
    const uuid = randomString(6);
    notifications.update(map => {
        map.set(uuid, {
            type: messageType,
            message: message,
            timeout: duration,
            title: title,
        })
        return map;
    })

    if (duration > 0) {
        setTimeout(() => {
            notifications.update(map => {
                map.delete(uuid);
                return map;
            })
        }, duration);
    }
};

export function remove(uuid: string) {
    notifications.update(map => {
        map.delete(uuid);
        return map;
    })
}

// gen random string
function randomString(length: number) {
    let result = '';
    const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    const charactersLength = characters.length;
    for (let i = 0; i < length; i++) {
        result += characters.charAt(Math.floor(Math.random() * charactersLength));
    }
    
    return result;
}
// import { writable, derived } from 'svelte/store';

/** Store for your data. 
This assumes the data you're pulling back will be an array.
If it's going to be an object, default this to an empty object.
**/
// export const apiData = writable([]);

/** Data transformation.
For our use case, we only care about the drink names, not the other information.
Here, we'll create a derived store to hold the drink names.
**/
// export const drinkNames = derived(apiData, ($apiData) => {
//   if ($apiData.drinks){
//     return $apiData.drinks.map(drink => drink.strDrink);
//   }
//   return [];
// });

// export const usersList = derived(apiData, ($apiData) => {
//  let d= $apiData.map(user => user.login);

//   return d;
//   // if ($apiData.drinks){
//   //   return $apiData.drinks.map(drink => drink.strDrink);
//   // }
//   // return [];
// });