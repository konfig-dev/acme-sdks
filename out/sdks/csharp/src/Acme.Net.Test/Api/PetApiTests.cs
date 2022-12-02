/*
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */

using System;
using System.IO;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.Reflection;
using RestSharp;
using NUnit.Framework;

using Acme.Net.Client;
using Acme.Net.Api;
using Acme.Net.Model;

namespace Acme.Net.Test
{
    /// <summary>
    ///  Class for testing PetApi
    /// </summary>
    /// <remarks>
    /// This file is automatically generated by OpenAPI Generator (https://openapi-generator.tech).
    /// Please update the test case below to test the API endpoint.
    /// </remarks>
    public class PetApiTests
    {
        private PetApi instance;

        /// <summary>
        /// Setup before each unit test
        /// </summary>
        [SetUp]
        public void Init()
        {
            instance = new PetApi();
        }

        /// <summary>
        /// Clean up after each unit test
        /// </summary>
        [TearDown]
        public void Cleanup()
        {

        }

        /// <summary>
        /// Test an instance of PetApi
        /// </summary>
        [Test]
        public void InstanceTest()
        {
            // TODO uncomment below to test 'IsInstanceOf' PetApi
            //Assert.IsInstanceOf(typeof(PetApi), instance);
        }

        
        /// <summary>
        /// Test AddPet
        /// </summary>
        [Test]
        public void AddPetTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //Pet pet = null;
            //var response = instance.AddPet(pet);
            //Assert.IsInstanceOf(typeof(Pet), response, "response is Pet");
        }
        
        /// <summary>
        /// Test DeletePet
        /// </summary>
        [Test]
        public void DeletePetTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //long petId = null;
            //string apiKey = null;
            //instance.DeletePet(petId, apiKey);
            
        }
        
        /// <summary>
        /// Test FindPetsByStatus
        /// </summary>
        [Test]
        public void FindPetsByStatusTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //List<string> status = null;
            //var response = instance.FindPetsByStatus(status);
            //Assert.IsInstanceOf(typeof(List<Pet>), response, "response is List<Pet>");
        }
        
        /// <summary>
        /// Test FindPetsByTags
        /// </summary>
        [Test]
        public void FindPetsByTagsTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //List<string> tags = null;
            //var response = instance.FindPetsByTags(tags);
            //Assert.IsInstanceOf(typeof(List<Pet>), response, "response is List<Pet>");
        }
        
        /// <summary>
        /// Test GetPetById
        /// </summary>
        [Test]
        public void GetPetByIdTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //long petId = null;
            //var response = instance.GetPetById(petId);
            //Assert.IsInstanceOf(typeof(Pet), response, "response is Pet");
        }
        
        /// <summary>
        /// Test UpdatePet
        /// </summary>
        [Test]
        public void UpdatePetTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //Pet pet = null;
            //var response = instance.UpdatePet(pet);
            //Assert.IsInstanceOf(typeof(Pet), response, "response is Pet");
        }
        
        /// <summary>
        /// Test UpdatePetWithForm
        /// </summary>
        [Test]
        public void UpdatePetWithFormTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //long petId = null;
            //string name = null;
            //string status = null;
            //instance.UpdatePetWithForm(petId, name, status);
            
        }
        
        /// <summary>
        /// Test UploadFile
        /// </summary>
        [Test]
        public void UploadFileTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //long petId = null;
            //string additionalMetadata = null;
            //System.IO.Stream file = null;
            //var response = instance.UploadFile(petId, additionalMetadata, file);
            //Assert.IsInstanceOf(typeof(ApiResponse), response, "response is ApiResponse");
        }
        
    }

}
