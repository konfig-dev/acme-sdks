/**
 *
 * Please note:
 * This class is auto generated by Konfig (https://konfigthis.com).
 * Do not edit this file manually.
 *
 */

@file:Suppress(
    "ArrayInDataClass",
    "EnumEntryName",
    "RemoveRedundantQualifierName",
    "UnusedImport"
)

package com.acme.client.models


import com.squareup.moshi.Json

/**
 * A tag for a pet
 *
 * @param id 
 * @param name 
 */


data class Tag (

    @Json(name = "id")
    val id: kotlin.Long? = null,

    @Json(name = "name")
    val name: kotlin.String? = null

)

