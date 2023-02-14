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

import com.acme.client.models.Category
import com.acme.client.models.Tag

import com.squareup.moshi.Json

/**
 * A pet for sale in the pet store
 *
 * @param name 
 * @param photoUrls 
 * @param tags 
 * @param id 
 * @param category 
 * @param status pet status in the store
 */


data class Pet (

    @Json(name = "name")
    val name: kotlin.String,

    @Json(name = "photoUrls")
    val photoUrls: kotlin.collections.List<kotlin.String>,

    @Json(name = "tags")
    val tags: kotlin.collections.List<Tag>? = null,

    @Json(name = "id")
    val id: kotlin.Long? = null,

    @Json(name = "category")
    val category: Category? = null,

    /* pet status in the store */
    @Json(name = "status")
    @Deprecated(message = "This property is deprecated.")
    val status: Pet.Status? = null

) {

    /**
     * pet status in the store
     *
     * Values: available,pending,sold
     */
    enum class Status(val value: kotlin.String) {
        @Json(name = "available") available("available"),
        @Json(name = "pending") pending("pending"),
        @Json(name = "sold") sold("sold");
    }
}

