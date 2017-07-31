/**
 * Created by sapelkinav on 31/07/2017.
 */
'use strict';

console.log('starting function');
exports.handle = function (e, ctx, cb) {
    console.log('processing event: %j', e);
    cb(null, {hello: 'world'});

};